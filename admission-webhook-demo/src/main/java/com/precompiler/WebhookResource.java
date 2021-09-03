package com.precompiler;

import io.fabric8.kubernetes.api.model.admission.v1.AdmissionReview;
import io.fabric8.kubernetes.api.model.admission.v1.AdmissionReviewBuilder;
import lombok.extern.slf4j.Slf4j;

import javax.ws.rs.Consumes;
import javax.ws.rs.POST;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Response;

/**
 * @author Richard Li
 */
@Slf4j
@Produces(MediaType.APPLICATION_JSON)
@Consumes(MediaType.APPLICATION_JSON)
@Path("/")
public class WebhookResource {
    @Path("/validate")
    @POST
    public Response validate(String payload) {
        log.info("validating request: {}", payload);

        AdmissionReview response = new AdmissionReviewBuilder().withNewResponse().withAllowed(Boolean.TRUE).withUid("DummyUID").endResponse().build();
        return Response.ok(response).build();
    }
}
