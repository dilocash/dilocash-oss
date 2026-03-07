/**
 * Copyright (c) 2026 dilocash
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */

import { createConnectTransport } from "@connectrpc/connect-web";

const BASE_URL = "http://localhost:8080";

export const getConnectTransport = (session: any) => {
  const transport = createConnectTransport({
      baseUrl: process.env.EXPO_PUBLIC_API_URL || BASE_URL,
      interceptors: [
        (next) => async (req) => {
          if (session?.access_token) {
            req.header.set(
              "Authorization",
              `Bearer ${session.access_token}`,
            );
          }
          return await next(req);
        },
      ],
  });
    return transport;
};

export default getConnectTransport;