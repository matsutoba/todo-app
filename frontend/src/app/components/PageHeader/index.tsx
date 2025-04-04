import { logoutAction } from "@/app/actions/login";
import { useActionState } from "react";

const PageHeader = () => {
  const [, formAction] = useActionState(logoutAction, { error: "" });

  const handleLogout = () => {
    if (confirm("ログアウトしますか？")) {
      formAction();
    }
  };

  return (
    <div className="flex justify-between items-center p-4">
      <h1 className="text-2xl font-bold">Task App</h1>
      <div>
        <form action={handleLogout}>
          <button
            type="submit"
            className="bg-blue-500 text-white px-2 py-1 rounded-md text-sm"
          >
            ログアウト
          </button>
        </form>
      </div>
    </div>
  );
};

export default PageHeader;
