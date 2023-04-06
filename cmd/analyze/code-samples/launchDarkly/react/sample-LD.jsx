import React from "react";

export const MyReactComponent = () => {
    const backgroundColorFlag = stringVariation("LD-string-flag-jsx", "green", environment)
    const backgroundSize = numberVariation("LD-number-flag-jsx", 16, environment)
    const showBtn = boolVariation("LD-bool-flag-jsx", true, environment)

    return (
        <div
            style={{
                height: "200px",
                width: "200px",
                backgroundColor: backgroundColorFlag.getValue(),
            }}
        >
            {"I'm a square with color=" + backgroundColorFlag.getValue()}
            <button style={{color: btnColorFlag.getValue()}}>Button</button>
        </div>
    );
};
