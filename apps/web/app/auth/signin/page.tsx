/**
 * Copyright (c) 2026 dilocash
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */
import type { Metadata } from "next";
import SigninClient from "./signin-client";

export const metadata: Metadata = {
  title: "Sign In",
  description: "Sign in to your dilocash account.",
};

export default function SigninPage() {
  return <SigninClient />;
}
