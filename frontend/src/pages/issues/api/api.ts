const getAllIssues = async () => {
  const response = await fetch("http://localhost:8080/issues");
  return response.json();
};

export { getAllIssues };