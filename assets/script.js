function deletes() {
    const myNode = document.getElementById("start")
    myNode.innerHTML = ""
}

function displays(a) {
    var x = document.getElementById(a);
    x.style.display = "block"
}

function displaysnot(a) {
    var x = document.getElementById(a);
    x.style.display = "none"
} 

// <script>
// input.addEventListener("keyup", function(event) {
// if (event.keyCode === 13) {
//     event.preventDefault();
//     document.getElementById("searchbtn").click();
// }
// })

// </script>