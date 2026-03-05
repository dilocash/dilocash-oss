/**
 * Copyright (c) 2026 dilocash
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */
"use client";

import { AuthForm } from "@dilocash/ui/components/auth/auth-form";
import { supabase } from "../../lib/supabase/client";
import { useEffect } from "react";
import { useAuth } from "@dilocash/ui/auth/provider";
import { useRouter } from "solito/navigation";

export default function LoginScreen() {
    const { session, isLoading } = useAuth()
    const { replace } = useRouter()
    useEffect(() => {
        // If session exists, redirect to main screen
        if (!isLoading && session) {
            replace('/main')
        }
    }, [session, isLoading, replace])
    console.log(isLoading)
    if (isLoading) return null

    return (
        <>
            <div>asgfasgasg</div>
            <AuthForm supabase={supabase} />
        </>
    );
}
