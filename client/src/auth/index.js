import Google from "next-auth/providers/google";
import GitHub from "next-auth/providers/github";

import NextAuth, { User, NextAuthConfig, Session } from "next-auth";
import Credentials from "next-auth/providers/credentials";

import { JWT } from "@auth/core/jwt";

export const BASE_PATH = "/api/auth";

export const { handlers, auth, signIn, signOut } = NextAuth({
  basePath: BASE_PATH,
  providers: [Google, GitHub],
  // session: {
  //   strategy: "database"
  // },
  pages: {
    signIn: "/login",
  },
  callbacks: {
    async jwt({ token, user }) {
      console.log("Entering JWT callback");
      if (user) {
        console.log("Found user");
        token.id = user.id;

        let response, data;
        try {
          response = await fetch(`${process.env.BACKEND_SERVER}/auth/signin`, {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify({
              id: user.id,
              email: user.email,
              name: user.name,
              image: user.image,
            }),
            credentials: "include",
          });
        } catch (err) {
          console.log("Could not fetch data from backend: " + err);
        }
        try {
          data = await response.json();
        } catch (err) {
          console.log("Could not deserialize json: " + err + ` ${response}`);
        }

        token.backendToken = data.token;
      }
      return token;
    },

    async session({ session, token }) {
      session.user.id = token.id;
      // let response, data;
      // try {
      //   response = await fetch(`${process.env.BACKEND_SERVER}/auth/signin`, {
      //     method: "POST",
      //     headers: {
      //       "Content-Type": "application/json",
      //     },
      //     body: JSON.stringify({
      //       id: session.user.id,
      //       email: session.user.email,
      //       name: session.user.name,
      //       image: session.user.image,
      //     }),
      //     credentials: "include",
      //   });
      // } catch (err) {
      //   console.log("Could not fetch data from backend: " + err);
      // }
      // try {
      //   data = await response.json();
      // } catch (err) {
      //   console.log("Could not deserialize json: " + err + ` ${response}`);
      // }

      session.user.backendToken = data.token;
      return session;
    },

    // async signIn({ user, account }) {
    //   console.log("Going to sign in at /api/v1/auth/signin");
    //   let response, data;
    //   try {
    //     response = await fetch(`${process.env.BACKEND_SERVER}/auth/signin`, {
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
    //     });
    //   } catch (err) {
    //     console.log("Could not fetch data from backend: " + err);
    //   }
    //   try {
    //     data = await response.json();
    //   } catch (err) {
    //     console.log("Could not deserialize json: " + err + ` ${response}`);
    //   }

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
