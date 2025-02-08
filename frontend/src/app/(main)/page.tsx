import { get } from "@/repository/todo";
import TaskEditButton from "../components/TaskEditButton";
import TaskDeleteButton from "../components/TaskDeleteButton";
import { MdAddTask } from "react-icons/md";
import Link from "next/link";

const MainPage = async () => {
  const todos = await get();
  return (
    <div>
      <Link
        href="/create"
        className="flex items-center gap-2 text-blue-500 hover:text-blue-600 justify-end mb-4"
      >
        <MdAddTask className="text-2xl cursor-pointer" />
        <span>タスクを作成</span>
      </Link>
      <table className="w-full">
        <thead className="bg-gray-200 text-left">
          <tr>
            <th className="p-2 w-2/4">タスク名</th>
            <th className="p-2 w-1/4">タスク内容</th>
            <th className="p-2 w-1/4">タスク完了</th>
            <th className="p-2 w-1/4"></th>
          </tr>
        </thead>
        <tbody>
          {todos.map((todo) => (
            <tr key={todo.id} className="border-b border-gray-200">
              <td className="p-2 w-1/3">{todo.title}</td>
              <td className="p-2 w-1/3">{todo.description}</td>
              <td className="p-2 w-1/3">
                {todo.completed ? "完了" : "未完了"}
              </td>
              <td className="p-2 w-1/4 flex gap-2">
                <TaskEditButton id={todo.id} />
                <TaskDeleteButton id={todo.id} />
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default MainPage;
