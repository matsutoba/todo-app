import { FaTrashAlt } from "react-icons/fa";

interface TaskDeleteButtonProps {
  id: string;
}

const TaskDeleteButton: React.FC<TaskDeleteButtonProps> = ({ id }) => {
  return (
    <form action="">
      <button type="submit">
        <FaTrashAlt className="hover:text-gray-400 text-lg cursor-pointer size-4" />
      </button>
    </form>
  );
};

export default TaskDeleteButton;
