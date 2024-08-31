import Google from "next-auth/providers/google"
import GitHub from "next-auth/providers/github"


import NextAuth, { User, NextAuthConfig } from "next-auth"
import Credentials from "next-auth/providers/credentials"

export const BASE_PATH = "/api/auth"

interface ExtendedUser extends User {
  sessionToken?: string;
}

export const { handlers, auth, signIn, signOut } = NextAuth({
  basePath: BASE_PATH,
  providers: [Google, GitHub],
  callbacks: {
    async jwt({ token, account }) {
      if (account) {
        token.accessToken = account.access_token
      }
      return token
    },
    // async signIn({ user, account }) {
    //   const data: { sessionToken: string } = await fetch("https://localhost:8001/api/v1/auth/signin", {
    //     method: "POST",
    //     body: JSON.stringify({ user, account })
    //   }).then((res) => res.json());

    //   if (data.sessionToken) {
    //     (user as ExtendedUser).sessionToken = data.sessionToken;
    //     return true;
    //   }

    //   return false;

    // }
    
    // async session({ session, token, user }) {
    //   session.user.id = token.id;
    
    //   return session
    // },
  },
})