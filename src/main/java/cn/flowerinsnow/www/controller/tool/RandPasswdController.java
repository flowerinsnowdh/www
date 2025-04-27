/*
 * www
 * Copyright (C) 2025  flowerinsnow
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */
package cn.flowerinsnow.www.controller.tool;

import org.springframework.http.HttpStatusCode;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.servlet.ModelAndView;

import java.security.SecureRandom;
import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

@Controller
public class RandPasswdController {
    private static final List<Character> NUMBERS = List.of('0', '1', '2', '3', '4', '5', '6', '7', '8', '9');
    private static final List<Character> LOWERCASES = List.of('a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z');
    private static final List<Character> UPPERCASES = List.of('A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z');
    private static final List<Character> CHARACTERS = List.of('~', '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '_', '+');
    private static final SecureRandom RANDOM = new SecureRandom();

    @GetMapping("/tool/randpasswd")
    public String getToolRandpasswd(Model model) {
        model.addAttribute("title", "随机密码生成器");
        model.addAttribute("length", 16);
        model.addAttribute("number", true);
        model.addAttribute("lowercase", true);
        model.addAttribute("uppercase", true);
        model.addAttribute("characters", true);
        return "tool/randpasswd";
    }

    @PostMapping("/tool/randpasswd")
    public ModelAndView postToolRandpasswd(
            @RequestParam(value = "length") int length,
            @RequestParam(value = "number", required = false) boolean useNumber,
            @RequestParam(value = "lowercase", required = false) boolean useLowercase,
            @RequestParam(value = "uppercase", required = false) boolean useUppercase,
            @RequestParam(value = "characters", required = false) boolean useCharacters
    ) {
        ModelAndView modelAndView = new ModelAndView("tool/randpasswd");
        modelAndView.addObject("title", "随机密码生成器");
        modelAndView.addObject("length", length);
        modelAndView.addObject("number", useNumber);
        modelAndView.addObject("lowercase", useLowercase);
        modelAndView.addObject("uppercase", useUppercase);
        modelAndView.addObject("characters", useCharacters);
        if (length <= 0 || length > 32) {
            modelAndView.addObject("errorMessages", Collections.singletonList("长度必须位于1-32之间"));
            modelAndView.setStatus(HttpStatusCode.valueOf(400));
            return modelAndView;
        } else {
            ArrayList<Character> list = new ArrayList<>();
            if (useNumber) {
                list.addAll(RandPasswdController.NUMBERS);
            }
            if (useLowercase) {
                list.addAll(RandPasswdController.LOWERCASES);
            }
            if (useUppercase) {
                list.addAll(RandPasswdController.UPPERCASES);
            }
            if (useCharacters) {
                list.addAll(RandPasswdController.CHARACTERS);
            }
            if (list.isEmpty()) {
                modelAndView.addObject("errorMessages", Collections.singletonList("咱至少要选一项吧？"));
                modelAndView.setStatus(HttpStatusCode.valueOf(400));
                return modelAndView;
            }
            char[] result = new char[length];
            for (int i = 0; i < length; i++) {
                result[i] = list.get(RandPasswdController.RANDOM.nextInt(list.size()));
            }
            modelAndView.addObject("result", new String(result));
        }
        return modelAndView;
    }
}
