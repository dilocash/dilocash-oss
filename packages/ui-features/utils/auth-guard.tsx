/**
 * Copyright (c) 2026 dilocash
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */

import { useEffect } from 'react'
import { useRouter } from 'solito/navigation'
import { useAuth } from './auth-provider'

export function AuthGuard({ children }: { children: React.ReactNode }) {
    const { user, isLoading } = useAuth()
    const { replace } = useRouter()
    useEffect(() => {
        if (!isLoading && !user) {
            replace('/')
        }
    }, [user, isLoading, replace])

    if (isLoading || !user) return null

    return <>{children}</>
}