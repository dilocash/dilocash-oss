"use client";
import "@dilocash/ui/global.css";  // Make sure this path is correct

import { GluestackUIProvider } from "@dilocash/ui";

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