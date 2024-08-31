import "./globals.css";
import { Inter } from "next/font/google";
import Header from "@/components/Header";
import Footer from "@/components/Footer";
import { NextAuthProvider } from "./session-provider";

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
          {/* <Header /> */}
          {children}
          {/* <Footer /> */}
        </NextAuthProvider>
      </body>
    </html>
  );
}
