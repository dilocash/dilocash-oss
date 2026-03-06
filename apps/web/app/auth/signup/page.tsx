/**
 * Copyright (c) 2026 dilocash
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */
"use client";

import { SignupForm } from "@dilocash/ui/components/auth/signup-form";
import { useEffect } from "react";
import { useAuth } from "@dilocash/ui/auth/provider";
import { useRouter } from "solito/navigation";

export default function SignupScreen() {
    const { session, isLoading } = useAuth()
    const { replace } = useRouter()
    useEffect(() => {
        // If session exists, redirect to main screen
        if (!isLoading && session) {
            replace('/main')
        }
    }, [session, isLoading, replace])
    if (isLoading) return null

    return (
        <SignupForm />
    );
}
