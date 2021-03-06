"use strict";

//TODO: Custom Errors and Loading Spinner
$(function () {
    $("#config_form").on("submit", function (e) {
        e.preventDefault();
        var username = $("#config_username").val();
        var password = $("#config_password").val();
        var confirm_password = $("#config_confirm_password").val();
        var remember_me = $("#config_remember_me").is(":checked");
        if (password !== confirm_password) {
            alert("Passwords Mismatch");
            return
        }

        $.ajax({
            url: "/v1/configure",
            data: JSON.stringify({
                username: username,
                password: password,
                remember_me: remember_me
            }),
            contentType: "application/json",
            type: "POST",
            success: function success() {
                location.href = "/v1/";
            },
            error: function error(data, status, err) {
                alert(err);
            }
        });
    });
});