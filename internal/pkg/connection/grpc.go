/**
 * Copyright 2020 Napptive
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package connection

import (
	"fmt"
	"github.com/napptive/catalog-cli/pkg/config"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

// GetConnection creates a connection with a gRPC server.
func GetConnection(cfg *config.ConnectionConfig) (*grpc.ClientConn, error) {
	return GetNonTLSConnection(cfg)
}

// GetNonTLSConnection returns a plain connection with the playground server.
func GetNonTLSConnection(cfg *config.ConnectionConfig) (*grpc.ClientConn, error) {
	log.Debug().Msg("using insecure connection with the Catalog-Manager")
	return grpc.Dial(fmt.Sprintf("%s:%d", cfg.CatalogAddress, cfg.CatalogPort), grpc.WithInsecure())
}
