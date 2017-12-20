+++
images = ["/assets/lstm/LSTM-cell.png"]
description = ""
categories = ["category"]
tags = ["tag1", "tag2"]
draft = true
title = "Parsing mathematical equation to generate computation graphs - Software 1.0/1.5/2.0 in go"
date = 2017-12-18T16:47:27+01:00
+++

In my last article, I have developped a recurrent neural network in pure go without any third party library.

As an example, I did an implementation of a character based generation (the famous Shakespeare example).
I have tried to tune the hyperparameters, but I haven't been able to reach a very usable text.

Actually, without any randomness in the generation process, the output was recurrent.

The point is that the toy I made is based on a vanilla RNN. And Vanillas RNNs are suffering from the [vanishing gradient problem](https://en.wikipedia.org/wiki/Vanishing_gradient_problem).
This is a well known problem, and one solution is to change the core model for a more robust network called __L__ong __S__hort __T__erm __M__emory network (LSTM for short).

# Implementing an LSTM

LSTM are a bit more complex than vanilla RNN. Therefore, a naive go implementation will be a harder.



## Side note about software 2.0

# Equations are graphs

## Gorgonia

# Good ol' software 1.0

## Lexer/Parser

### goyacc

{{< highlight go >}}
// Forward pass as described here https://en.wikipedia.org/wiki/Long_short-term_memory#LSTM_with_a_forget_gate
func (l *lstm) fwd(inputVector, prevHidden, prevCell *G.Node) (hidden, cell *G.Node) {
        // Helper function for clarity
        set := func(ident, equation string) *G.Node {
                res, _ := l.parser.Parse(equation)
                l.parser.Set(ident, res)
                return res 
        } 

        l.parser.Set(`xₜ`, inputVector)
        l.parser.Set(`hₜ₋₁`, prevHidden)
        l.parser.Set(`cₜ₋₁`, prevCell)
        set(`iₜ`, `σ(Wᵢ·xₜ+Uᵢ·hₜ₋₁+Bᵢ)`)
        set(`fₜ`, `σ(Wf·xₜ+Uf·hₜ₋₁+Bf)`) // dot product made with ctrl+k . M
        set(`oₜ`, `σ(Wₒ·xₜ+Uₒ·hₜ₋₁+Bₒ)`)
        // ċₜis a vector of new candidates value
        set(`ĉₜ`, `tanh(Wc·xₜ+Uc·hₜ₋₁+Bc)`) // c made with ctrl+k c >
        ct := set(`cₜ`, `fₜ*cₜ₋₁+iₜ*ĉₜ`)
        set(`hc`, `tanh(cₜ)`)
        ht, _ := l.parser.Parse(`oₜ*hc`)
        return ht, ct
}
{{</ highlight >}}

If you don't have the correct font to display the unicode character, you may find a picture [here](/assets/lstm/uni-code.png)

![image](/assets/lstm/LSTM.png)

# Conclusion

