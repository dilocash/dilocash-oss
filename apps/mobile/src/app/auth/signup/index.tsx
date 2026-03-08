/**
 * Copyright (c) 2026 dilocash
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */
"use client";

import supabase from "../../lib/supabase/client";
import { SignupFormOTP } from "@dilocash/ui/components/auth/signup-form-otp";
import { VerifyCodeForm } from "@dilocash/ui/components/auth/verify-code-form";
import { useState } from "react";
import { useRouter } from "solito/navigation";
import { Box } from "@dilocash/ui/components/ui/box";

export default function SignupScreen() {
    const { replace } = useRouter()
    const [emailForVerification, setEmailForVerification] = useState(null);
    const handleOtpSent = (email: any) => {
        console.debug(`otp sent to ${email}`)
        setEmailForVerification(email);
    };

    const handleOtpVerified = (email: any) => {
        console.debug(`otp verified for ${email}`)
        setEmailForVerification(null);
        replace('/', {
            experimental: {
                nativeBehavior: 'stack-replace',
                isNestedNavigator: false,
            },
        })
    };

    return (
        <Box className="h-full">
            {!emailForVerification ? (
                <SignupFormOTP supabase={supabase} onOTPSent={handleOtpSent} />
            ) : (
                <VerifyCodeForm supabase={supabase} email={emailForVerification} onOTPVerified={handleOtpVerified} />
            )}
        </Box>
    );
}
