import React from "react";
const fetchUserData =async () => {
  const query_string=`
	query{  
  user {
	    id
      name
	  }
  }
  `
  const settings = {
    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body:JSON.stringify({query:query_string})
  };
  try{
    const response=await fetch("http://localhost:8080/query",settings)
    console.log(response)
    const data=await response.json()
    console.log(data)
  }
  catch(error)
  {
    console.log(error)
  }
};
function App() {
  return (
    <div>
      <button onClick={fetchUserData}>Show User </button>
    </div>
  );
}

export default App;
