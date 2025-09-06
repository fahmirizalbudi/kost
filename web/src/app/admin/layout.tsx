import type { Metadata } from "next";
import SideNavigation from "./components/SideNavigation";
import { menuSideNavigation } from "../data/menu";

export const metadata: Metadata = {
  title: "Kostopia - Admin Dashboard",
  description: "App that used for rent an dorm",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body>
        <SideNavigation menu={menuSideNavigation} />
        {children}
      </body>
    </html>
  );
}
