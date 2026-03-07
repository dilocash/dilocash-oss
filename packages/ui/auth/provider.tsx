import { createContext, useContext, useEffect, useState } from 'react'
import { Session, User, SupabaseClient } from '@supabase/supabase-js'

type AuthContextType = {
    supabase: SupabaseClient | null
    session: Session | null
    user: User | null
    isLoading: boolean
}

const AuthContext = createContext<AuthContextType>({
    supabase: null,
    session: null,
    user: null,
    isLoading: true,
})

export function AuthProvider({ supabase, children }: { supabase: SupabaseClient, children: React.ReactNode }) {
    const [session, setSession] = useState<Session | null>(null)
    const [isLoading, setIsLoading] = useState(true)

    useEffect(() => {
        // get session
        supabase.auth.getSession().then(({ data: { session } }) => {
            console.debug('getSession', session)
            setSession(session)
            setIsLoading(false)
        })

        // listen for changes (login, logout, etc.)
        const { data: { subscription } } = supabase.auth.onAuthStateChange((_event, session) => {
            console.debug('onAuthStateChange', _event, session)
            setSession(session)
            setIsLoading(false)
        })

        return () => subscription.unsubscribe()
    }, [])

    return (
        <AuthContext.Provider value={{ session, supabase, user: session?.user ?? null, isLoading }}>
            {children}
        </AuthContext.Provider>
    )
}

// Hook to handle auth state
export const useAuth = () => useContext(AuthContext)