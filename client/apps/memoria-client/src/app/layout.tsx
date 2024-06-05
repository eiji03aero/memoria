import type { Metadata } from 'next';
import { Inter } from 'next/font/google';

import { Providers } from '@/Providers';
import '@/app/globals.css';

const inter = Inter({ subsets: ['latin'] });

export const metadata: Metadata = {
  title: 'Memoria',
  description: 'A frinedly app helps you cherish your moments.',
};
export const fetchCache = 'default-no-store';
export const dynamic = 'force-dynamic';

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <head>
        <meta name="application-name" content="Memoria" />
        <meta name="apple-mobile-web-app-capable" content="yes" />
        <meta name="apple-mobile-web-app-status-bar-style" content="default" />
        <meta name="apple-mobile-web-app-title" content="Memoria" />
        <meta name="description" content="Best app to store memories" />
        <meta name="format-detection" content="telephone=no" />
        <meta name="mobile-web-app-capable" content="yes" />

        <link rel="apple-touch-icon" href="/icons/apple-touch-icon.png" />

        <link rel="icon" type="image/png" sizes="32x32" href="/icons/favicon-32x32.png" />
        <link rel="icon" type="image/png" sizes="16x16" href="/icons/favicon-16x16.png" />
        <link rel="manifest" href="/manifest.json" />
        <link rel="shortcut icon" href="/icons/favicon.ico" />

        <meta name="twitter:card" content="summary" />
        <meta name="twitter:url" content="https://memoria-app.com" />
        <meta name="twitter:title" content="Memoria" />
        <meta name="twitter:description" content="Best app to store memories" />
        <meta name="twitter:image" content="https://memoria.com/icons/android-chrome-192x192.png" />
        <meta name="twitter:creator" content="@EijiOsakabe" />
        <meta property="og:type" content="website" />
        <meta property="og:title" content="Memoria" />
        <meta property="og:description" content="Best app to store memories" />
        <meta property="og:site_name" content="Memoria" />
        <meta property="og:url" content="https://memoria-app.com" />
        <meta property="og:image" content="https://memoria-app.com/icons/apple-touch-icon.png" />
      </head>
      <body className={inter.className}>
        <Providers>{children}</Providers>
      </body>
    </html>
  );
}
