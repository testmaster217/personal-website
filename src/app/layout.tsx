import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  // TODO: Add more metadata.
  title: "Collin Vesel",
  description: "My personal website. Includes some basic information about me and I hope to add (many) other features soon. :)",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body>
        {children}
      </body>
    </html>
  );
}
