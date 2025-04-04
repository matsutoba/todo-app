import Link from "next/link";
import { FaPen } from "react-icons/fa";

interface TaskEditButtonProps {
  id: number;
}

const TaskEditButton: React.FC<TaskEditButtonProps> = ({ id }) => {
  return (
    <Link href={`/edit/${id}`}>
      <FaPen className="hover:text-gray-400 text-lg cursor-pointer size-4" />
    </Link>
  );
};

export default TaskEditButton;
