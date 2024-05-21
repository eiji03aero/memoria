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
      <body className={inter.className}>
        <Providers>{children}</Providers>
      </body>
    </html>
  );
}
