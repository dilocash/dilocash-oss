import { useEffect } from 'react'
import { useRouter } from 'solito/navigation'
import { useAuth } from './provider'

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