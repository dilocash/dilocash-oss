/**
 * Copyright (c) 2026 dilocash
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */
"use client";

import { SigninForm } from "@dilocash/ui/components/auth/signin-form";
import supabase from "../../lib/supabase/client";
import { useEffect } from "react";
import { useAuth } from "@dilocash/ui/auth/provider";
import { useRouter } from "solito/navigation";

export default function SigninScreen() {
    const { session, isLoading } = useAuth()
    const { push } = useRouter()
    useEffect(() => {
        // If session exists, redirect to main screen
        if (!isLoading && session) {
            push('/main')
        }
    }, [session, isLoading])
    if (isLoading) return null

    return (
        <SigninForm supabase={supabase} />
    );

}
