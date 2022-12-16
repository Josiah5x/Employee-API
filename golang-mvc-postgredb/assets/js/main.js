function GetSingleUser() {
    var txt;
    var person = prompt("Please enter your name:");
    if (person == null || person == "") {
      txt = "User cancelled the prompt.";
    } else {
      txt = "Hello " + person + "! How are you today?";
    }
      
      window.location.assign(person)
    // <a href="http://localhost:8080/v1/employee/02502d23-2df2-4c9f-a5ad-7bd7b99cecc3"></a>
  }

  // function UpdateEmployee() {
  //   var txt;
  //   var id = prompt("Please enter your username:");

  //   window.location.assign(id)

  // }

