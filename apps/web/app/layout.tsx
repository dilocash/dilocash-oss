"use client";
import "../styles/global.css";
import "../styles/styles.css";

import { GluestackUIProvider } from "@dilocash/ui/components/ui/gluestack-ui-provider";
import i18n, { initI18n } from '@dilocash/i18n';
import { useEffect, useState } from 'react';
import { AppLoader } from '@dilocash/ui/components/app-loader';

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const [isReady, setIsReady] = useState(false);
  useEffect(() => {
    // init i18n
    
  const setup = async () => {
    await initI18n(true);
    //await initRxDB(); // init IndexedDB
    var millisecondsToWait = 50000;
    setTimeout(function() {
      // Whatever you want to do after the wait
      setIsReady(true);
    }, millisecondsToWait);
  };
  setup();

  }, []);
  return (
    <html lang={i18n.language}>
      <body>
        <GluestackUIProvider mode="light">
          {!isReady ? <AppLoader subMessage="..." isWeb={true} /> : children}
        </GluestackUIProvider>
      </body>
    </html>
  );
}