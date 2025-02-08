import TaskEditForm from "@/app/components/TaskEditForm.tsx";
import { getById } from "@/repository/todo";

interface Params {
  params: {
    id: string;
  };
}

const EditTaskPage = async ({ params }: Params) => {
  const { id } = await params;
  const todo = await getById(Number(id));

  return <TaskEditForm todo={todo} />;
};

export default EditTaskPage;
