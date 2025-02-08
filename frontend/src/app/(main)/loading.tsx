const Loading = () => {
  return (
    <div
      className="h-full flex justify-center items-center"
      aria-label="loading"
    >
      <div className="animate-spin h-10 w-10 border-4 border-gray-400 rounded-full border-t-transparent"></div>
    </div>
  );
};

export default Loading;
