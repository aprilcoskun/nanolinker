"use strict";

//TODO: Custom Errors and Loading Spinner
$(function () {
    $("#sign_in_form").on("submit", function (e) {
        e.preventDefault();
        var username = $("#sign_in_username").val();
        var password = $("#sign_in_password").val();
        var remember_me = $("#sign_in_remember_me").prop(":checked");
        $.ajax({
            url: "/v1/login",
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