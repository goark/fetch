package fetch

import "errors"

var (
	// ErrNullPointer indicates a nil receiver or nil embedded value access.
	ErrNullPointer = errors.New("null reference instance")
	// ErrInvalidRequest indicates a request construction or execution failure.
	ErrInvalidRequest = errors.New("invalid HTTP request")
	// ErrInvalidURL indicates an invalid URL input.
	ErrInvalidURL = errors.New("invalid URL")
	// ErrHTTPStatus indicates non-success HTTP status.
	ErrHTTPStatus = errors.New("bad HTTP status")
)

/* Copyright 2021-2026 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
