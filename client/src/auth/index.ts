import Google from "next-auth/providers/google"
import GitHub from "next-auth/providers/github"


import NextAuth, { User, NextAuthConfig, Session } from "next-auth";
import Credentials from "next-auth/providers/credentials";

import { JWT } from "@auth/core/jwt";

export const BASE_PATH = "/api/auth";

interface ExtendedUser extends User {
  backendToken?: string;
}

export interface CustomSession extends Session {
  user: ExtendedUser;
}

interface BackendJWT extends JWT {
  id?: string;
  backendToken?: string;
}

export const { handlers, auth, signIn, signOut } = NextAuth({
  basePath: BASE_PATH,
  providers: [Google, GitHub],
  pages: {
    signIn: "/login",
  },
  callbacks: {
    async jwt({ token, user }: { token: JWT; user: ExtendedUser }) {
      if (user) {
        // const response = await fetch(
        //   `${process.env.BACKEND_SERVER}auth/signin`,
        //   {
        //     method: "POST",
        //     headers: {
        //       "Content-Type": "application/json",
        //     },
        //     body: JSON.stringify({ user }),
        //   }
        // );
        // const data = await response.json();
        // token.backendToken = data.token;
        token.id = user.id;
      }
      return token;
    },

    async session({ session, token }: { session: CustomSession; token: JWT }) {
      session.user.id = token.id as string;
      // session.user.backendToken = token.backendToken;
      return session;
    },

    // async signIn({ user, account }) {
    //   console.log("Going to sign in at /api/v1/auth/signin");
    //   const response = await fetch(
    //     `${process.env.BACKEND_SERVER}/auth/signin`,
    //     {
    //       method: "POST",
    //       headers: {
    //         "Content-Type": "application/json",
    //       },
    //       body: JSON.stringify({
    //         id: user.id,
    //         email: user.email,
    //         name: user.name,
    //         image: user.image,
    //       }),
    //       credentials: "include",
    //     }
    // //   );

    //   // If the API call was successful, allow sign in
    //   if (response.ok) {
    //     return true;
    //   } else {
    //     // If there was an error storing the user details, you might want to prevent sign in
    //     return false;
    //   }
    // },
  },
});