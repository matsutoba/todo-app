"use client";

import { useActionState } from "react";
import { FormState, loginAction } from "../actions/login";

const LoginPage = () => {
  const initialState: FormState = {
    error: "",
  };
  const [state, formAction, isPending] = useActionState(
    loginAction,
    initialState
  );

  return (
    <div className="h-screen flex flex-col items-center justify-center">
      <h1 className="text-2xl font-bold">ログイン</h1>
      <form action={formAction}>
        <div>
          {state.error && (
            <p className="text-red-500 flex justify-center mt-4">
              {state.error}
            </p>
          )}
          <div className="grid grid-cols-2 gap-4 mt-8">
            <label htmlFor="email">ユーザID</label>
            <input
              type="email"
              id="email"
              name="email"
              className="border border-gray-300 rounded-md p-2"
            />
          </div>
          <div className="grid grid-cols-2 gap-4 mt-4">
            <label htmlFor="password">パスワード</label>
            <input
              type="password"
              id="password"
              name="password"
              className="border border-gray-300 rounded-md p-2 disabled:opacity-50"
              disabled={isPending}
            />
          </div>
        </div>
        <button
          type="submit"
          className="bg-blue-500 text-white p-2 rounded-md mt-4 flex justify-center w-full"
        >
          ログイン
        </button>
      </form>
    </div>
  );
};

export default LoginPage;
