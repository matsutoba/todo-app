import { get } from "@/repository/todo";
import TaskEditButton from "../components/TaskEditButton";
import TaskDeleteButton from "../components/TaskDeleteButton";

const MainPage = async () => {
  const todos = await get();
  console.log(todos);

  return (
    <div>
      <h1>タスク一覧</h1>
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
