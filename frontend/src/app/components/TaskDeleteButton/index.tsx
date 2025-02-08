"use client";

import { FaTrashAlt } from "react-icons/fa";
import { useActionState, useState, useTransition } from "react";
import { deleteTodo } from "@/app/actions/todo";
import ConfirmModal from "../ConfirmModal";
interface TaskDeleteButtonProps {
  id: number;
}

const TaskDeleteButton: React.FC<TaskDeleteButtonProps> = ({ id }) => {
  const [isOpen, setIsOpen] = useState(false);
  const [, startTransitionq] = useTransition();

  const deleteTodoWithId = deleteTodo.bind(null, id);
  const [, formAction, isPending] = useActionState(deleteTodoWithId, {
    error: "",
  });

  const handleDelete = async () => {
    startTransitionq(async () => {
      await formAction();
    });
  };

  return (
    <>
      <button
        disabled={isPending}
        onClick={() => setIsOpen(true)}
        className="hover:text-gray-400 text-lg cursor-pointer size-4 disabled:opacity-50"
      >
        <FaTrashAlt />
      </button>
      {isOpen && (
        <ConfirmModal
          title="タスクを削除しますか？"
          message="この操作は元に戻すことができません。"
          isOpen={isOpen}
          onCancel={() => setIsOpen(false)}
          onConfirm={() => {
            setIsOpen(false);
            handleDelete();
          }}
        />
      )}
    </>
  );
};

export default TaskDeleteButton;
