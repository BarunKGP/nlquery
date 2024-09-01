import "./globals.css";
import { Inter } from "next/font/google";
import { NextAuthProvider } from "./session-provider";
import { ThemeProvider } from "./ThemeProvider";

const inter = Inter({ subsets: ["latin"] });

export const metadata = {
  title: "nlQuery",
  description: "Talk to your charts",
};

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body className={`${inter.className}`}>
        <NextAuthProvider>
          <ThemeProvider>{children}</ThemeProvider>
        </NextAuthProvider>
      </body>
    </html>
  );
}
