"use client";
import "../styles/global.css";

import { GluestackUIProvider } from "@dilocash/ui/components/ui/gluestack-ui-provider";
import i18n, { initI18n } from '@dilocash/i18n';
import { useEffect, useState } from 'react';

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const [isReady, setIsReady] = useState(false);
  useEffect(() => {
    // Inicializamos indicando que es Web para activar el detector
    initI18n(true).then(() => setIsReady(true));
  }, []);
  if (!isReady) return <html>
      <body>
      </body>
    </html>; // Or a loading spinner
  return (
    <html lang={i18n.language}>
      <body>
        <GluestackUIProvider mode="light">
          {children}
        </GluestackUIProvider>
      </body>
    </html>
  );
}