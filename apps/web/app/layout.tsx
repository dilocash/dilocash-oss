/**
 * Copyright (c) 2026 dilocash
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */

import "../styles/global.css";
import "../styles/styles.css";
import type { Viewport, Metadata } from 'next';
import { headers } from 'next/headers';
import ClientLayout from './client-layout';
import { SerwistProvider } from "./serwist/serwist";

const APP_NAME = "Dilocash";
const APP_DEFAULT_TITLE = "Dilocash";
const APP_TITLE_TEMPLATE = "%s | Dilocash";
const APP_DESCRIPTION = "Manage your cash and expenses easily with dilocash.";

export const viewport: Viewport = {
  width: 'device-width',
  initialScale: 1,
  maximumScale: 1,
  userScalable: false,
  themeColor: '#ffffff',
};

export const metadata: Metadata = {
  title: {
    default: APP_DEFAULT_TITLE,
    template: APP_TITLE_TEMPLATE,
  },
  description: APP_DESCRIPTION,
  applicationName: APP_NAME,
  authors: [{ name: 'dilocash' }],
  keywords: ['cash', 'expenses', 'finance', 'management'],
  appleWebApp: {
    capable: true,
    statusBarStyle: 'default',
    title: 'dilocash',
  },
  formatDetection: {
    telephone: false,
  },
  openGraph: {
    type: 'website',
    siteName: APP_NAME,
    title: {
      default: APP_DEFAULT_TITLE,
      template: APP_TITLE_TEMPLATE,
    },
    description: APP_DESCRIPTION,
  },
  twitter: {
    card: 'summary',
    title: {
      default: APP_DEFAULT_TITLE,
      template: APP_TITLE_TEMPLATE,
    },
    description: APP_DESCRIPTION,
  },
};

const SUPPORTED_LOCALES = ['es', 'en'] as const;
type Locale = (typeof SUPPORTED_LOCALES)[number];
const DEFAULT_LOCALE: Locale = 'en';

/**
 * Parses the Accept-Language header and returns the best matching
 * supported locale, falling back to DEFAULT_LOCALE.
 *
 * e.g. "es-419,es;q=0.9,en;q=0.8" → "es"
 */
function detectLocale(acceptLanguage: string | null): Locale {
  if (!acceptLanguage) return DEFAULT_LOCALE;

  const preferred = acceptLanguage
    .split(',')
    .map((part) => {
      const [tag, q] = part.trim().split(';q=');
      return { lang: tag?.split('-')[0]?.toLowerCase() ?? '', q: q ? parseFloat(q) : 1 };
    })
    .sort((a, b) => b.q - a.q)
    .map((p) => p.lang);

  return (preferred.find((l): l is Locale => (SUPPORTED_LOCALES as readonly string[]).includes(l))) ?? DEFAULT_LOCALE;
}

export default async function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const headersList = await headers();
  const locale = detectLocale(headersList.get('accept-language'));
  return (
    <html lang={locale} dir="ltr">
      <head />
      <body>
        <SerwistProvider swUrl="/sw.js" disable={process.env.NODE_ENV === "development"}>
          <ClientLayout locale={locale}>{children}</ClientLayout>
        </SerwistProvider>
      </body>
    </html>
  );
}
