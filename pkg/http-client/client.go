// Copyright 2024 The Archivista Contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/in-toto/archivista/pkg/api"
	"github.com/in-toto/go-witness/dsse"
)

type ArchivistaClient struct {
	BaseURL    string
	GraphQLURL string
	*http.Client
}

func CreateArchivistaClient(httpClient *http.Client, baseURL string) (*ArchivistaClient, error) {
	client := ArchivistaClient{
		BaseURL: baseURL,
		Client:  http.DefaultClient,
	}
	if httpClient != nil {
		client.Client = httpClient
	}
	var err error
	client.GraphQLURL, err = url.JoinPath(client.BaseURL, "query")
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (ac *ArchivistaClient) DownloadDSSE(ctx context.Context, gitoid string) (dsse.Envelope, error) {
	reader, err := api.DownloadReadCloserWithHTTPClient(ctx, ac.Client, ac.BaseURL, gitoid)
	if err != nil {
		return dsse.Envelope{}, err
	}
	env := dsse.Envelope{}
	if err := json.NewDecoder(reader).Decode(&env); err != nil {
		return dsse.Envelope{}, err
	}
	return env, nil
}

func (ac *ArchivistaClient) DownloadReadCloser(ctx context.Context, gitoid string) (io.ReadCloser, error) {
	return api.DownloadReadCloserWithHTTPClient(ctx, ac.Client, ac.BaseURL, gitoid)
}

func (ac *ArchivistaClient) DownloadWithWriter(ctx context.Context, gitoid string, dst io.Writer) error {
	return api.DownloadWithWriterWithHTTPClient(ctx, ac.Client, ac.BaseURL, gitoid, dst)
}

func (ac *ArchivistaClient) Store(ctx context.Context, envelope dsse.Envelope) (api.StoreResponse, error) {
	return api.Store(ctx, ac.BaseURL, envelope)
}

func (ac *ArchivistaClient) StoreWithReader(ctx context.Context, r io.Reader) (api.StoreResponse, error) {
	return api.StoreWithReader(ctx, ac.BaseURL, r)
}

type GraphQLRequestBodyInterface struct {
	Query     string      `json:"query"`
	Variables interface{} `json:"variables,omitempty"`
}

type GraphQLResponseInterface struct {
	Data   interface{}
	Errors []api.GraphQLError `json:"errors,omitempty"`
}

func (ac *ArchivistaClient) GraphQLRetrieveSubjectResults(
	ctx context.Context,
	gitoid string,
) (api.RetrieveSubjectResults, error) {
	return api.GraphQlQuery[api.RetrieveSubjectResults](
		ctx,
		ac.BaseURL,
		api.RetrieveSubjectsQuery,
		api.RetrieveSubjectVars{Gitoid: gitoid},
	)
}

func (ac *ArchivistaClient) GraphQLRetrieveSearchResults(
	ctx context.Context,
	algo string,
	digest string,
) (api.SearchResults, error) {
	return api.GraphQlQuery[api.SearchResults](
		ctx,
		ac.BaseURL,
		api.SearchQuery,
		api.SearchVars{Algorithm: algo, Digest: digest},
	)
}

func (ac *ArchivistaClient) GraphQLQueryIface(
	ctx context.Context,
	query string,
	variables interface{},
) (*GraphQLResponseInterface, error) {
	reader, err := ac.GraphQLQueryReadCloser(ctx, query, variables)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	gqlRes := GraphQLResponseInterface{}
	dec := json.NewDecoder(reader)
	if err := dec.Decode(&gqlRes); err != nil {
		return nil, err
	}
	if len(gqlRes.Errors) > 0 {
		return nil, fmt.Errorf("graph ql query failed: %v", gqlRes.Errors)
	}
	return &gqlRes, nil
}

func (ac *ArchivistaClient) GraphQLQueryToDst(ctx context.Context, query string, variables interface{}, dst interface{}) error {
	reader, err := ac.GraphQLQueryReadCloser(ctx, query, variables)
	if err != nil {
		return err
	}
	defer reader.Close()
	dec := json.NewDecoder(reader)
	if err := dec.Decode(&dst); err != nil {
		return err
	}
	return nil
}

func (ac *ArchivistaClient) GraphQLQueryReadCloser(
	ctx context.Context,
	query string,
	variables interface{},
) (io.ReadCloser, error) {
	requestBodyMap := GraphQLRequestBodyInterface{
		Query:     query,
		Variables: variables,
	}
	requestBodyJSON, err := json.Marshal(requestBodyMap)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, ac.GraphQLURL, bytes.NewReader(requestBodyJSON))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := ac.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		defer res.Body.Close()
		errMsg, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(errMsg))
	}
	return res.Body, nil
}
