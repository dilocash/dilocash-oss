/**
 * Copyright (c) 2026 dilocash
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */
"use client";

import { AuthProvider } from "@dilocash/ui/auth/provider";
import { supabase } from "../lib/supabase/client";
import MainScreen from "./main/page";

export default function Home() {
  return (
    <AuthProvider supabase={supabase}>
      <MainScreen />
    </AuthProvider>
  );
}
