import type { Metadata } from "next";
import localFont from "next/font/local";
import "./styles/globals.scss";

const sfProText = localFont({
  src: [
    { path: '/fonts/SFProText-Regular.ttf', weight: '400', style: 'normal' },
    { path: '/fonts/SFProText-RegularItalic.ttf', weight: '400', style: 'italic' },

    { path: '/fonts/SFProText-Medium.ttf', weight: '500', style: 'normal' },
    { path: '/fonts/SFProText-MediumItalic.ttf', weight: '500', style: 'italic' },

    { path: '/fonts/SFProText-Semibold.ttf', weight: '600', style: 'normal' },
    { path: '/fonts/SFProText-SemiboldItalic.ttf', weight: '600', style: 'italic' },

    { path: '/fonts/SFProText-Bold.ttf', weight: '700', style: 'normal' },
    { path: '/fonts/SFProText-BoldItalic.ttf', weight: '700', style: 'italic' },

    { path: '/fonts/SFProText-Heavy.ttf', weight: '900', style: 'normal' },
    { path: '/fonts/SFProText-HeavyItalic.ttf', weight: '900', style: 'italic' },
  ],
  variable: '--font-sfpro',
});

export const metadata: Metadata = {
  title: "Kostopia - Rent a Dorm",
  description: "App that used for rent an dorm",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={`${sfProText.variable}`}>
        {children}
      </body>
    </html>
  );
}
