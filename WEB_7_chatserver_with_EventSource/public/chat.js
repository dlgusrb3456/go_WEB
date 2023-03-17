$(function(){
    if (!window.EventSource){
        alert("No event source")
        return
    }

    var $chatlog = $('#chat-log')
    var $chatmsg = $('#chat-msg')

    var isBlank = function(string){
        return string == null || string.trim() === "";
    };
    var username;
    while (isBlank(username)){
        username = prompt("What's your name?");
        if (!isBlank(username)){
            $('#user-name').html('<b>'+username+'</b>')
        }
    }

    $('#input-form').on('submit',function(e){ //id input-form에서 submit이라는 event가 발생하면 funciton()을 수행해라
        $.post('/messages',{
            msg: $chatmsg.val(),
            name: username
        }); //$로 시작하면 jquery를 사용하겠다는 것임. jquery로 post를 /message url로 보내라.
        $chatmsg.val("")
        $chatmsg.focus()
        return false;
    });

    var addMessage = function(data){
        var text = ""
        if (!isBlank(data.name)){
            text = '<strong>'+data.name+': </strong> ';
        }
        text += data.msg
        $chatlog.prepend('<div><span>'+text + '</span></div>');
    };

    var es = new EventSource('/stream')
    es.onopen = function(e){
        $.post('users/',{
            name: username
        });
    }
    es.onmessage = function(e){
        var msg = JSON.parse(e.data)
        addMessage(msg)
    }

    window.onbeforeunload = function(){
        $.ajax({
            url: "/users?username=" + username,
            type: "DELETE"
        });
        es.close()
    };
})