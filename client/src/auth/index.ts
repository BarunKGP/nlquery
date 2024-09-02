import Google from "next-auth/providers/google"
import GitHub from "next-auth/providers/github"


import NextAuth, { User, NextAuthConfig } from "next-auth"
import Credentials from "next-auth/providers/credentials"

import { DefaultSession } from "next-auth";

export const BASE_PATH = "/api/auth";

interface ExtendedUser extends User {
  sessionToken?: string;
}

interface CustomSession extends DefaultSession {
  accessToken?: string;
}

export const { handlers, auth, signIn, signOut } = NextAuth({
  basePath: BASE_PATH,
  providers: [Google, GitHub],
  callbacks: {
    async jwt({ token, account }) {
      if (account) {
        token.accessToken = account.access_token;
      }
      return token;
    },

    async signIn({ user, account, profile }) {
      const response = await fetch(
        `${process.env.BACKEND_SERVER}/api/v1/auth/login`,
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            userId: user.id,
            email: user.email,
            name: user.name,
            // Add any other relevant user details
          }),
        }
      );

      // If the API call was successful, allow sign in
      if (response.ok) {
        return true;
      } else {
        // If there was an error storing the user details, you might want to prevent sign in
        return false;
      }
    },

    async session({ session, token }): Promise<CustomSession> {
      return {
        ...session,
        accessToken: token.accessToken as string,
      };
    },
  },
});