
<form action="javascript:void(0);" class="ui form">
    <h2 class="title"><i class="user icon"></i>User Information</h2>
        {{template "best_practice/users/form/body.tpl" .}}
    <input class="ui primary button" id="save" type="submit" value="Create" onclick="SaveUser();">
    <input class="ui button" id="cancel" type="submit" value="Cancel" onclick="CancelUser();">
</form>


<script type="text/javascript">

    function SaveUser() {
        var userName = document.getElementById("username").value;
        var email = document.getElementById("email").value;
        var name = document.getElementById("name").value;
        $.ajax({
            async: false,
            type: "post",
            dataType: "json",
            url: "/users/saveUser",
            data: { "username": userName, "email": email, "name": name },
            success: function (data) {
                console.log("user saved data = " + data);
                $('body').html(data);
            }
        });
    }

    function ChangePassword() {
        window.alert("reset password pending");
    }
</script>