"use strict";

$(function () {
    $('#new_expires_at')
        .calendar({
            minDate: new Date()
        });
});

function openEditRow(id) {
    var idId = "#table_id_" + id;
    var urlId = "#table_url_" + id;
    var oldId = $("#table_id_" + id).html();
    var oldUrl = $("#table_url_" + id).html();
    $(idId).html(toInput(idId, oldId));
    $(urlId).html(toInput(urlId, oldUrl));
    $("#table_actions_" + id).html(cancelOrSaveButtons(id))
}

function cancelEditRow(id) {
    var idVal = $("#table_id_" + id + "_input").val();
    var urlVal = $("#table_url_" + id + "_input").val();
    $("#table_id_" + id).html(idVal);
    $("#table_url_" + id).html(urlVal);
    $("#table_actions_" + id).html(editOrDeleteButtons(id));
}

function saveEditLink(id) {
}

function saveLink() {
    var newId = $("#new_id").val();
    var newUrl = $("#new_url").val();
    console.log(newId, newUrl);
    $.ajax({
        url: "/v1/link",
        data: JSON.stringify({id: newId, url: newUrl}),
        contentType: "application/json",
        type: "POST",
        success: function success() {
            location.href = "/v1";
        },
        error: function error(data, status, err) {
            alert(err);
        }
    });
}

function deleteRow(id) {
    $('.basic.modal').modal('setting', 'closable', false).modal('show');
}

function toInput(id, val) {
    return '<div class="ui input small fluid"><input id="' + id.substr(1) + '_input" type="text" value="' + val + '"/></div>';
}

function cancelOrSaveButtons(id) {
    return '<div class="ui buttons">' +
        '<button class="ui green inverted button"  onclick="saveEditLink(\'' + id + '\')"><i class="check icon"></i>Save</button>' +
        '<button class="ui red  inverted button" onclick="cancelEditRow(\'' + id + '\')"><i class="close icon"></i>Cancel</button>' +
        '</div>';
}

function editOrDeleteButtons(id) {
    return '<div class="ui buttons">' +
        '<button class="ui icon olive button" onclick="openEditRow(\'' + id + '\')"><i class="pen icon"></i>Edit</button>' +
        '<button class="ui icon red button" onclick="deleteRow(\'' + id + '\')"><i class="trash alternate icon"></i>Delete</button>' +
        '</div>';
}