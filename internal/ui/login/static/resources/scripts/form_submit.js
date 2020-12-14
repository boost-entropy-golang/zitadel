function disableSubmit(checks, button) {
    let form = document.getElementsByTagName('form')[0];
    let inputs = form.getElementsByTagName('input');
    button.disabled = true;
    addRequiredEventListener(inputs, checks, form, button);
    disableDoubleSubmit(form, button);
}

function addRequiredEventListener(inputs, checks, form, button) {
    let eventType = 'input';
    for (i = 0; i < inputs.length; i++) {
        if (inputs[i].required) {
            eventType = 'input';
            if (inputs[i].type === 'checkbox') {
                eventType = 'click';
            }
            inputs[i].addEventListener(eventType, function () {
                toggleButton(checks, form, inputs, button);
            });
        }
    }
}

function disableDoubleSubmit(form, button) {
    form.addEventListener('submit', function() {
        document.body.classList.add('waiting');
        button.disabled = true;
    })
}

function toggleButton(checks, form, inputs, button) {
    if (checks !== undefined) {
        if (checks() === false) {
            button.disabled = true;
            return
        }
    }
    button.disabled = !allRequiredDone(form, inputs);
}

function allRequiredDone(form, inputs) {
    for (i = 0; i < inputs.length; i++) {
        if (inputs[i].required) {
            if (inputs[i].type === 'checkbox' && !inputs[i].checked) {
                return false
            }
            if (inputs[i].value === '') {
                return false
            }
        }
    }
    return true;
}