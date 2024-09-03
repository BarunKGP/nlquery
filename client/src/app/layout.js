import "./globals.css";
import { Inter } from "next/font/google";
import { NextAuthProvider } from "./session-provider";
import { ThemeProvider } from "./ThemeProvider";
import Head from "next/head";
import Link from "next/link";

const inter = Inter({ subsets: ["latin"] });

export const metadata = {
  title: "nlQuery",
  description: "Talk to your charts",
};

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <Head>
        <Link
          rel="apple-touch-icon"
          type="image/png"
          href={"/apple-touch-icon.png"}
          aria-setsize={180}
        />
        <Link
          rel="icon"
          type="image/png"
          href={"/favicon-32x32.png"}
          aria-setsize={32}
        />
        <Link
          rel="icon"
          type="image/png"
          href={"/favicon-16x16.png"}
          aria-setsize={16}
        />
        <Link rel="manifest" type="image/png" href={"/site.webmanifest"} />
      </Head>
      <body className={`${inter.className}`}>
        <NextAuthProvider>
          <ThemeProvider>{children}</ThemeProvider>
        </NextAuthProvider>
      </body>
    </html>
  );
}
