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
package cn.flowerinsnow.www.controller.exception;

import cn.flowerinsnow.www.message.error.badrequest.RequestParameterMissing;
import cn.flowerinsnow.www.message.error.badrequest.RequestParameterTypeMismatch;
import cn.flowerinsnow.www.util.ControllerUtil;
import org.springframework.http.HttpStatus;
import org.springframework.ui.Model;
import org.springframework.web.bind.MissingServletRequestParameterException;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.method.annotation.MethodArgumentTypeMismatchException;

@ControllerAdvice
public class ExceptionHandlerController {
    @ResponseStatus(HttpStatus.BAD_REQUEST)
    @ExceptionHandler(MissingServletRequestParameterException.class)
    public String onBadRequest(MissingServletRequestParameterException exception, Model model) {
        model.addAttribute("title", "400 Bad Request");
        ControllerUtil.addErrorMessages(model, RequestParameterMissing.ofName(exception.getParameterName()));
        return "error/400";
    }

    @ResponseStatus(HttpStatus.BAD_REQUEST)
    @ExceptionHandler(MethodArgumentTypeMismatchException.class)
    public String onBadRequest(MethodArgumentTypeMismatchException exception, Model model) {
        model.addAttribute("title", "400 Bad Request");
        ControllerUtil.addErrorMessages(model, RequestParameterTypeMismatch.ofName(exception.getName()));
        return "error/400";
    }
}
