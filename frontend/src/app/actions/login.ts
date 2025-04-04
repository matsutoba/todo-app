"use server";

import { cookies } from "next/headers";
import { redirect } from "next/navigation";

export interface FormState {
  error: string;
}

export const loginAction = async (state: FormState, formData: FormData) => {
  const email = formData.get("email");
  const password = formData.get("password");

  const response = await fetch(`${process.env.API_URL}/users/login`, {
    method: "POST",
    body: JSON.stringify({ email, password }),
    headers: {
      "Content-Type": "application/json",
    },
  });

  if (!response.ok) {
    const s = { ...state, error: "ユーザIDまたはパスワードが違います" };
    return s;
  }

  const json = await response.json();
  const token = json.token;

  (await cookies()).set("token", token, { httpOnly: true, secure: true });

  redirect("/");
};

export const logoutAction = async () => {
  (await cookies()).delete("token");
  redirect("/login");
};
