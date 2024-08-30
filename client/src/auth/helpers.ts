"use server"
import { signIn as naSignIn, signOut as naSignOut } from "."

export async function signIn(signInOptions: any) {
    await naSignIn(signInOptions);
}

export async function signOut(signOutOptions: any) {
    await naSignOut(signOutOptions);
}