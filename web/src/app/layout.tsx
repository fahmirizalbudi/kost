import type { Metadata } from "next";
import localFont from "next/font/local";
import "./styles/globals.scss";

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
      <body className={``}>
        {children}
      </body>
    </html>
  );
}
