import React from 'react';

const InputText = (props: { placeholder: string, onChange: (target: string) => void, defaultValue: string }) => {
    return (
        <input
            type={'text'}
            placeholder={props.placeholder ? props.placeholder : ''}
            onChange={e => props.onChange(e.target.value)}
            value = {props.defaultValue}
        />
    )
}

export default InputText;