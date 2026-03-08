/**
 * Copyright (c) 2026 dilocash
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */

"use client";

import { GluestackUIProvider } from "@dilocash/ui-components/components/ui/gluestack-ui-provider";
import DatabaseProvider from '../lib/database-provider';
import { initI18n } from '@dilocash/i18n';
import { useEffect, useState } from 'react';
import { AppLoader } from '@dilocash/ui-components/components/app-loader';
import { AuthProvider } from "@dilocash/ui-features/utils/auth-provider";
import supabase from "../lib/supabase/client";

export default function ClientLayout({
  children,
  locale,
}: Readonly<{
  children: React.ReactNode;
  /** Locale resolved server-side from Accept-Language header */
  locale: string;
}>) {
  const [isReady, setIsReady] = useState(false);

  useEffect(() => {
    const setup = async () => {
      // Seed i18n with the locale detected server-side so it doesn't
      // have to re-detect (and potentially flash to a different language).
      await initI18n(true, locale);
      setTimeout(() => setIsReady(true), 50);
    };
    setup();
  }, [locale]);

  return (
    <DatabaseProvider>
      <GluestackUIProvider mode="light">
        <AuthProvider supabase={supabase}>
          {!isReady ? <AppLoader subMessage="..." isWeb={true} /> : children}
        </AuthProvider>
      </GluestackUIProvider>
    </DatabaseProvider>
  );
}
