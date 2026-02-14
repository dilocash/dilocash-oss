"use client";
import "../styles/global.css"; 

import { GluestackUIProvider } from "@dilocash/ui/components/ui/gluestack-ui-provider";

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body>
        <GluestackUIProvider mode="light">
          {children}
        </GluestackUIProvider>
      </body>
    </html>
  );
}