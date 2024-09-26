"use server";
import { SignInOptions, SignInAuthorizationParams } from "next-auth/react";
import { signIn as naSignIn, signOut as naSignOut } from ".";
import { BuiltInProviderType } from "next-auth/providers";

type SignInOptionsType = {
  provider: BuiltInProviderType;
  options: SignInOptions;
  authorizationParams: SignInAuthorizationParams;
};

export async function signIn(signInOptions?: SignInOptionsType) {
  await naSignIn(
    signInOptions.provider,
    signInOptions.options,
    signInOptions.authorizationParams
  )
    .then((data) => {
      fetch(`${process.env.SERVER}/auth/signin`, {
        method: "POST",
        body: JSON.stringify(data, null, 2),
      });
    })
    .catch((err) => {
      console.log(`Signin error: ${JSON.stringify(err, null, 2)}`);
    });
}

